import { patchState, signalStore, withMethods, withState } from '@ngrx/signals';
import { BlogPost } from '../../models/blog-post.model';
import { inject } from '@angular/core';
import { BlogService } from '../../services/blog.service';

type BlogState = {
  posts: BlogPost[];
  loading: boolean;
  error: string | null;
  page: number;
  allLoaded: boolean;
};

const initialState: BlogState = {
  posts: [],
  loading: false,
  error: null,
  page: 1,
  allLoaded: false,
};

export const BlogStore = signalStore(
  { providedIn: 'root' },
  withState(initialState),
  withMethods((store, blogService = inject(BlogService)) => ({
    async loadMore() {
      if (store.loading() || store.allLoaded()) return;

      patchState(store, { loading: true });
      try {
        const newPosts = await blogService.getAll(store.page(), 10);
        if (newPosts.length < 10) {
          patchState(store, { allLoaded: true });
        }
        patchState(store, {
          posts: [...store.posts(), ...newPosts],
          page: store.page() + 1,
          loading: false,
        });
      } catch (error) {
        patchState(store, { error: 'Failed to fetch blog posts', loading: false });
      }
    },
  }))
);
