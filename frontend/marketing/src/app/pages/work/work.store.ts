import { patchState, signalStore, withMethods, withState } from '@ngrx/signals';
import { WorkContent } from '../../models/work.model';
import { inject } from '@angular/core';
import { WorkService } from '../../services/work.service';

type WorkState = {
  content: WorkContent[];
  loading: boolean;
  error: string | null;
  page: number;
  allLoaded: boolean;
};

const initialState: WorkState = {
  content: [],
  loading: false,
  error: null,
  page: 1,
  allLoaded: false,
};

export const WorkStore = signalStore(
  { providedIn: 'root' },
  withState(initialState),
  withMethods((store, workService = inject(WorkService)) => ({
    async loadMore() {
      if (store.loading() || store.allLoaded()) return;

      const pageSize = 6;

      patchState(store, { loading: true });
      try {
        const newContent: WorkContent[] = await workService.getAll(store.page(), pageSize);

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
        patchState(store, { error: `Failed to fetch work content. \nMessage: ${JSON.stringify(error)}`, loading: false });
      }
    },
  }))
);