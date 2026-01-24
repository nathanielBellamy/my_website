export interface HomeContent {
  id: string;
  title: string;
  content: string;
}

export interface HomeResponse {
  content: HomeContent[];
}