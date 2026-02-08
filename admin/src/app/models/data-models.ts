export interface Author {
  id: string;
  name: string;
}

export interface Tag {
  id: string;
  name: string;
}

export interface BlogPost {
  id: string;
  title: string;
  content: string;
  author: Author | null; // Use null for optional relationships
  tags: Tag[];
  createdAt: string; // ISO 8601 string
  updatedAt: string; // ISO 8601 string
}

export interface HomeContent {
  id: string;
  title: string;
  content: string;
}

export interface GrooveJrContent {
  id: string;
  title: string;
  content: string;
}

export interface AboutContent {
  id: string;
  title: string;
  content: string;
}

export interface TrackerData {
  ip: string;
}
