import { patchState, signalStore, withMethods, withState } from '@ngrx/signals';
import { GrooveJrContent } from '../../models/groove-jr.model';
import { inject } from '@angular/core';
import { GrooveJrService } from '../../services/groove-jr.service';

type GrooveJrState = {
  content: GrooveJrContent[];
  loading: boolean;
  error: string | null;
  page: number;
  allLoaded: boolean;
};

const initialState: GrooveJrState = {
  content: [],
  loading: false,
  error: null,
  page: 1,
  allLoaded: false,
};

export const GrooveJrStore = signalStore(
  { providedIn: 'root' },
  withState(initialState),
  withMethods((store, grooveJrService = inject(GrooveJrService)) => ({
    async loadMore() {
      if (store.loading() || store.allLoaded()) return;

      patchState(store, { loading: true });
      try {
        const newContent = await grooveJrService.getAll(store.page(), 10);
        if (newContent.length < 10) {
          patchState(store, { allLoaded: true });
        }
        patchState(store, {
          content: [...store.content(), ...newContent],
          page: store.page() + 1,
          loading: false,
          error: null,
        });
      } catch (error) {
        patchState(store, { error: 'Failed to fetch groove-jr content', loading: false });
      }
    },
  }))
);