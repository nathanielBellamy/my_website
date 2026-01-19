import { patchState, signalStore, withMethods, withState } from '@ngrx/signals';
import { BlogPost } from '../../models/blog.model';
import { inject } from '@angular/core';
import { BlogService } from '../../services/blog.service';

type BlogState = {
  posts: BlogPost[];
  loading: boolean;
  error: string | null;
};

const initialState: BlogState = {
  posts: [],
  loading: false,
  error: null,
};

export const BlogStore = signalStore(
  { providedIn: 'root' },
  withState(initialState),
  withMethods((store, blogService = inject(BlogService)) => ({
    async loadPosts() {
      patchState(store, { loading: true });
      try {
        const response = await blogService.getAll();
        patchState(store, { posts: response.data, loading: false });
      } catch (error) {
        patchState(store, { error: 'Failed to fetch blog posts', loading: false });
      }
    },
  }))
);
