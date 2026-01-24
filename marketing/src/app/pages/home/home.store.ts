import { patchState, signalStore, withMethods, withState } from '@ngrx/signals';
import { HomeContent } from '../../models/home.model';
import { inject } from '@angular/core';
import { HomeService } from '../../services/home.service';

type HomeState = {
  content: HomeContent[];
  loading: boolean;
  error: string | null;
  allLoaded: boolean;
};

const initialState: HomeState = {
  content: [],
  loading: false,
  error: null,
  allLoaded: false,
};

export const HomeStore = signalStore(
  { providedIn: 'root' },
  withState(initialState),
  withMethods((store, homeService = inject(HomeService)) => ({
    async loadMore() {
      if (store.loading() || store.allLoaded()) return;

      const pageSize = 6;
      const currentPage = Math.ceil(store.content().length / pageSize) + 1;

      patchState(store, { loading: true });
      try {
        const response = await homeService.getAll(currentPage, pageSize);
        const newContent = response.content;

        patchState(store, {
          content: [...store.content(), ...newContent],
          loading: false,
        });
      } catch (error) {
        patchState(store, { error: `Failed to fetch home content. \nMessage: ${error}`, loading: false });
      }
    },
  }))
);