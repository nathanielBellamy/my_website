export interface Author {
  id: string;
  name: string;
}

export interface Tag {
  id: string;
  name: string;
  usageCount?: number;
}

export interface BlogPost {
  id: string;
  title: string;
  content: string;
  author: Author;
  tags: Tag[];
  createdAt: string;
  updatedAt: string;
  activatedAt?: string;
  deactivatedAt?: string;
  order: number;
}
