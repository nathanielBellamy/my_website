import { patchState, signalStore, withMethods, withState } from '@ngrx/signals';
import { HomeContent } from '../../models/home.model';
import { inject } from '@angular/core';
import { HomeService } from '../../services/home.service';

type HomeState = {
  content: HomeContent[];
  loading: boolean;
  error: string | null;
};

const initialState: HomeState = {
  content: [],
  loading: false,
  error: null,
};

export const HomeStore = signalStore(
  { providedIn: 'root' },
  withState(initialState),
  withMethods((store, homeService = inject(HomeService)) => ({
    async loadContent() {
      patchState(store, { loading: true });
      try {
        const response = await homeService.getAll();
        patchState(store, { content: response.data, loading: false });
      } catch (error) {
        patchState(store, { error: 'Failed to fetch home content', loading: false });
      }
    },
  }))
);
