#!/bin/bash
set -e

# Usage: ./deploy-image.sh <USER> <HOST> <SSH_KEY_PATH> <TARGET_ENV>

SSH_USER=$1
SSH_HOST=$2
SSH_KEY_PATH=$3
TARGET_ENV=$4 # "develop" or "production"

if [ -z "$SSH_USER" ] || [ -z "$SSH_HOST" ] || [ -z "$SSH_KEY_PATH" ] || [ -z "$TARGET_ENV" ]; then
  echo "Usage: $0 <USER> <HOST> <SSH_KEY_PATH> <TARGET_ENV>"
  exit 1
fi

echo "--------------------------------------------------"
echo "🚀 Starting Deployment to $TARGET_ENV"
echo "Target: $SSH_USER@$SSH_HOST"
echo "--------------------------------------------------"

# Ensure strict host key checking is off for CI/CD automation
# We use a custom ssh command wrapper for convenience
SSH_CMD="ssh -o StrictHostKeyChecking=no -i $SSH_KEY_PATH"
SCP_CMD="scp -o StrictHostKeyChecking=no -i $SSH_KEY_PATH"

# 1. Save and compress the Docker image
echo "📦 Saving and compressing Docker image (my_website_backend:latest)..."
# We pipe directly to gzip to save disk I/O
docker save my_website_backend:latest | gzip > backend.tar.gz
echo "✅ Image compressed to backend.tar.gz"

# 2. Transfer files
echo "📤 Transferring files to server..."
$SCP_CMD backend.tar.gz $SSH_USER@$SSH_HOST:/tmp/backend.tar.gz
$SCP_CMD docker-compose.prod.yml $SSH_USER@$SSH_HOST:~/docker-compose.yml
echo "   Securing .env file..."
$SSH_CMD $SSH_USER@$SSH_HOST "if [ -f ~/.env ]; then rm ~/.env; fi && mkdir -p ~/.env && chmod 700 ~/.env"
if [ -f .env/.env.production ]; then
  $SCP_CMD .env/.env.production $SSH_USER@$SSH_HOST:~/.env/.env.production
  $SSH_CMD $SSH_USER@$SSH_HOST "chmod 600 ~/.env/.env.production || true"
fi
# Also transfer the database init script if needed, though usually it's in the volume or image
# For this setup, we mount ./database/init.sql, so we must ensure it exists on remote
echo "   Ensuring database directory exists..."
$SSH_CMD $SSH_USER@$SSH_HOST "mkdir -p ~/database"
$SCP_CMD database/init.sql $SSH_USER@$SSH_HOST:~/database/init.sql

echo "   Transferring lifecycle scripts..."
$SSH_CMD $SSH_USER@$SSH_HOST "mkdir -p ~/lifecycle"
$SCP_CMD nixos/lifecycle/*.sh $SSH_USER@$SSH_HOST:~/lifecycle/
$SSH_CMD $SSH_USER@$SSH_HOST "chmod +x ~/lifecycle/*.sh"

echo "   Securing .env directory..."
$SSH_CMD $SSH_USER@$SSH_HOST "chmod 700 ~/.env || true"
echo "✅ Files transferred."

# 3. Execute deployment on server
echo "⚡ Executing remote deployment commands..."
$SSH_CMD $SSH_USER@$SSH_HOST << EOF
  set -e
  
  # Load new image
  echo "   [Remote] Loading Docker image..."
  gunzip -c /tmp/backend.tar.gz | docker load
  rm /tmp/backend.tar.gz

  # Backup logs
  mkdir -p logs
  echo "   [Remote] Backing up logs..."
  # We use || true so deployment doesn't fail if no container exists yet
  docker logs my_website_backend > logs/backend-\$(date +%Y%m%d-%H%M).log 2>&1 || true

  # Restart service
  echo "   [Remote] Restarting services..."
  # Export env vars so docker-compose can interpolate them (e.g. POSTGRES_USER, DATABASE_URL)
  # We can't use 'source' directly because .env files may have unquoted spaces in values.
  # Instead, read line-by-line and export safely.
  while IFS='=' read -r key value; do
    # Skip blank lines and comments
    [ -z "\$key" ] && continue
    case "\$key" in \#*) continue ;; esac
    export "\$key=\$value"
  done < ~/.env/.env.production
  # We use up -d which intelligently recreates containers only if image changed or config changed
  # We force recreate the backend to ensure it picks up the new image even if 'latest' tag confusion exists
  docker compose up -d --force-recreate backend || {
    echo "❌ Deployment failed! Fetching postgres-db logs for debugging:"
    docker logs my_website_db
    exit 1
  }
  docker compose up -d postgres-db # Ensure DB is up (won't restart if healthy)

  # Cleanup
  echo "   [Remote] Pruning old images..."
  docker image prune -f
EOF

echo "--------------------------------------------------"
echo "✅ Deployment to $TARGET_ENV complete!"
echo "--------------------------------------------------"

# Cleanup local artifact
rm backend.tar.gz
