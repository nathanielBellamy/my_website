import { patchState, signalStore, withMethods, withState } from '@ngrx/signals';
import { GrooveJrContent } from '../../models/groove-jr.model';
import { inject } from '@angular/core';
import { GrooveJrService } from '../../services/groove-jr.service';

type GrooveJrState = {
  content: GrooveJrContent[];
  loading: boolean;
  error: string | null;
};

const initialState: GrooveJrState = {
  content: [],
  loading: false,
  error: null,
};

export const GrooveJrStore = signalStore(
  { providedIn: 'root' },
  withState(initialState),
  withMethods((store, grooveJrService = inject(GrooveJrService)) => ({
    async loadContent() {
      patchState(store, { loading: true });
      try {
        const response = await grooveJrService.getAll();
        patchState(store, { content: response.data, loading: false });
      } catch (error) {
        patchState(store, { error: 'Failed to fetch GrooveJr content', loading: false });
      }
    },
  }))
);
