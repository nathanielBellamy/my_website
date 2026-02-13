export interface Author {
  id: string;
  name: string;
  activatedAt?: string | null;
  deactivatedAt?: string | null;
}

export interface Tag {
  id: string;
  name: string;
  activatedAt?: string | null;
  deactivatedAt?: string | null;
}

export interface BlogPost {
  id: string;
  title: string;
  content: string;
  author: Author | null; // Use null for optional relationships
  tags: Tag[];
  createdAt: string; // ISO 8601 string
  updatedAt: string; // ISO 8601 string
  activatedAt?: string | null;
  deactivatedAt?: string | null;
}

export interface HomeContent {
  id: string;
  title: string;
  content: string;
  activatedAt?: string | null;
  deactivatedAt?: string | null;
}

export interface GrooveJrContent {
  id: string;
  title: string;
  content: string;
  activatedAt?: string | null;
  deactivatedAt?: string | null;
}

export interface AboutContent {
  id: string;
  title: string;
  content: string;
  activatedAt?: string | null;
  deactivatedAt?: string | null;
}

export interface TrackerData {
  ip: string;
}
