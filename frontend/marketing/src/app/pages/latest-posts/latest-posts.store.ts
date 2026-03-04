import { patchState, signalStore, withMethods, withState } from '@ngrx/signals';
import { HomeContent } from '../../models/home.model';
import { inject } from '@angular/core';
import { HomeService } from '../../services/home.service';

type LatestPostsState = {
  content: HomeContent[];
  loading: boolean;
  error: string | null;
  page: number;
  allLoaded: boolean;
};

const initialState: LatestPostsState = {
  content: [],
  loading: false,
  error: null,
  page: 1,
  allLoaded: false,
};

export const LatestPostsStore = signalStore(
  { providedIn: 'root' },
  withState(initialState),
  withMethods((store, homeService = inject(HomeService)) => ({
    async loadMore() {
      if (store.loading() || store.allLoaded()) return;

      const pageSize = 6;

      patchState(store, { loading: true });
      try {
        const newContent: HomeContent[] = await homeService.getAll(store.page(), pageSize);

        if (newContent.length < pageSize) {
          patchState(store, { allLoaded: true });
        }

        patchState(store, {
          content: [...store.content(), ...newContent],
          page: store.page() + 1,
          loading: false,
          error: null,
        });
      } catch (error) {
        patchState(store, { error: `Failed to fetch latest posts content. \nMessage: ${error}`, loading: false });
      }
    },
  }))
);