import { patchState, signalStore, withMethods, withState } from '@ngrx/signals';
import { HomeContent } from '../../models/home.model';
import { inject } from '@angular/core';
import { HomeService } from '../../services/home.service';

type HomeState = {
  content: HomeContent[];
  loading: boolean;
  error: string | null;
  page: number;
  allLoaded: boolean;
};

const initialState: HomeState = {
  content: [],
  loading: false,
  error: null,
  page: 1,
  allLoaded: false,
};

export const HomeStore = signalStore(
  { providedIn: 'root' },
  withState(initialState),
  withMethods((store, homeService = inject(HomeService)) => ({
    async loadMore() {
      if (store.loading() || store.allLoaded()) return;

      patchState(store, { loading: true });
      try {
        const newContent = await homeService.getAll(store.page(), 10);
        if (newContent.length < 10) {
          patchState(store, { allLoaded: true });
        }
        patchState(store, {
          content: [...store.content(), ...newContent],
          page: store.page() + 1,
          loading: false,
        });
      } catch (error) {
        patchState(store, { error: 'Failed to fetch home content', loading: false });
      }
    },
  }))
);