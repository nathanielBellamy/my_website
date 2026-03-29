#!/bin/bash
set -e

# Load environment variables
if [ -f .env ]; then
  set -a
  source .env
  set +a
fi
if [ -f .env/.env.localhost ]; then
  set -a
  source .env/.env.localhost
  set +a
fi

# Configuration
DB_CONTAINER="my_website_db"
DB_USER="postgres"
DB_NAME="mw_db"
BACKUP_DIR="backups"
TIMESTAMP=$(date +"%Y-%m-%d_%H-%M-%S")
FILENAME="db_backup_${TIMESTAMP}.sql"
FILEPATH="${BACKUP_DIR}/${FILENAME}"

# Create backups directory if it doesn't exist
mkdir -p "$BACKUP_DIR"

echo "Creating database backup..."
if docker exec -e PGPASSWORD=postgres "$DB_CONTAINER" pg_dump -U "$DB_USER" "$DB_NAME" > "$FILEPATH"; then
    echo "Backup created successfully: $FILEPATH"
else
    echo "Error: Failed to create database backup."
    rm -f "$FILEPATH"
    exit 1
fi

echo "Compressing backup..."
gzip "$FILEPATH"
FILEPATH="${FILEPATH}.gz"
FILENAME="${FILENAME}.gz"
echo "Backup compressed: $FILEPATH"

echo "Sending email..."

python3 - <<EOF
import smtplib
import os
import sys
from email.mime.multipart import MIMEMultipart
from email.mime.base import MIMEBase
from email.mime.text import MIMEText
from email import encoders

smtp_host = os.environ.get('SMTP_HOST')
smtp_port = os.environ.get('SMTP_PORT')
smtp_user = os.environ.get('SMTP_USER')
smtp_pass = os.environ.get('SMTP_PASS')
filepath = "$FILEPATH"
filename = "$FILENAME"

if not all([smtp_host, smtp_port, smtp_user, smtp_pass]):
    print("Error: SMTP configuration missing in environment variables.")
    sys.exit(1)

try:
    msg = MIMEMultipart()
    msg['From'] = smtp_user
    msg['To'] = smtp_user
    msg['Subject'] = f"Database Backup: {filename}"
    
    body = "Database backup attached."
    msg.attach(MIMEText(body, 'plain'))
    
    with open(filepath, "rb") as attachment:
        part = MIMEBase("application", "octet-stream")
        part.set_payload(attachment.read())
    
    encoders.encode_base64(part)
    part.add_header(
        "Content-Disposition",
        f"attachment; filename= {filename}",
    )
    msg.attach(part)
    
    server = smtplib.SMTP(smtp_host, int(smtp_port))
    server.starttls()
    server.login(smtp_user, smtp_pass)
    text = msg.as_string()
    server.sendmail(smtp_user, smtp_user, text)
    server.quit()
    print(f"Email sent successfully to {smtp_user}")

except Exception as e:
    print(f"Error sending email: {e}")
    sys.exit(1)
EOF

echo "Backup process completed."
