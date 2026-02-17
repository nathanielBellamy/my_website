import { patchState, signalStore, withMethods, withState } from '@ngrx/signals';
import { BlogPost, Tag } from '../../models/blog-post.model';
import { inject } from '@angular/core';
import { BlogService } from '../../services/blog.service';

type BlogState = {
  posts: BlogPost[];
  loading: boolean;
  error: string | null;
  page: number;
  allLoaded: boolean;
  availableTags: Tag[];
  selectedTags: string[];
};

const initialState: BlogState = {
  posts: [],
  loading: false,
  error: null,
  page: 1,
  allLoaded: false,
  availableTags: [],
  selectedTags: [],
};

export const BlogStore = signalStore(
  { providedIn: 'root' },
  withState(initialState),
  withMethods((store, blogService = inject(BlogService)) => ({
    async loadMore() {
      if (store.loading() || store.allLoaded()) return;

      patchState(store, { loading: true });
      try {
        const newPosts = await blogService.getAll(store.page(), 10, store.selectedTags());
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
    async loadTags() {
      try {
        const tags = await blogService.getTags();
        patchState(store, { availableTags: tags });
      } catch (error) {
        console.error('Failed to fetch tags', error);
      }
    },
    async searchTags(query: string) {
        try {
            const tags = await blogService.getTags(query);
            patchState(store, { availableTags: tags });
        } catch (error) {
            console.error('Failed to search tags', error);
        }
    },
    async toggleTag(tagId: string) {
        const current = store.selectedTags();
        const newTags = current.includes(tagId) 
            ? current.filter(id => id !== tagId)
            : [...current, tagId];
        
        patchState(store, { 
            selectedTags: newTags,
            posts: [],
            page: 1,
            allLoaded: false
        });
        
        // Trigger loadMore to fetch the first page with new filters
        // We can't call loadMore() directly from here easily if we want to reuse the logic without duplicating it or extracting it.
        // Actually we can just call the service and patch state similar to loadMore, or extract a private helper if possible (not easy in withMethods object literal).
        // Or we can define loadMore as a method and call it using `this.loadMore()` ? No, `this` context is tricky here.
        // But we can recursively call the method exposed in the store... wait, the store object is not fully constructed yet.
        // Standard pattern is to extract the fetching logic or just duplicate the small part for now. 
        // Or better: just call blogService.getAll directly here.
        
        patchState(store, { loading: true });
        try {
            const newPosts = await blogService.getAll(1, 10, newTags);
            if (newPosts.length < 10) {
              patchState(store, { allLoaded: true });
            }
            patchState(store, {
              posts: newPosts,
              page: 2,
              loading: false,
            });
        } catch (error) {
            patchState(store, { error: 'Failed to fetch blog posts', loading: false });
        }
    }
  }))
);
