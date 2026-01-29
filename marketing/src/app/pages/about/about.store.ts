import { inject } from '@angular/core';
import { patchState, signalStore, withMethods, withState } from '@ngrx/signals';
import { AboutContent } from '../../models/about.model';
import { AboutService } from '../../services/about.service';

type AboutState = {
  content: AboutContent[];
  loading: boolean;
  error: string | null;
  page: number;
  allLoaded: boolean;
};

const initialState: AboutState = {
  content: [],
  loading: false,
  error: null,
  page: 1,
  allLoaded: false,
};

export const AboutStore = signalStore(
  { providedIn: 'root' },
  withState(initialState),
  withMethods((store, aboutService = inject(AboutService)) => ({
    async loadMore() {
      if (store.loading() || store.allLoaded()) return;

      patchState(store, { loading: true });
      try {
        const newContent = await aboutService.getAll(store.page(), 10);
        if (newContent.length < 10) {
          patchState(store, { allLoaded: true });
        }
        patchState(store, {
          content: [...store.content(), ...newContent],
          page: store.page() + 1,
          loading: false,
        });
      } catch (error) {
        patchState(store, { error: 'Failed to fetch about content', loading: false });
      }
    },
  }))
);