import { patchState, signalStore, withMethods, withState } from '@ngrx/signals';
import { AboutContent } from '../../models/about.model';
import { inject } from '@angular/core';
import { AboutService } from '../../services/about.service';

type AboutState = {
  content: AboutContent[];
  loading: boolean;
  error: string | null;
};

const initialState: AboutState = {
  content: [],
  loading: false,
  error: null,
};

export const AboutStore = signalStore(
  { providedIn: 'root' },
  withState(initialState),
  withMethods((store, aboutService = inject(AboutService)) => ({
    async loadContent() {
      patchState(store, { loading: true });
      try {
        const response = await aboutService.getAll();
        patchState(store, { content: response.data, loading: false });
      } catch (error) {
        patchState(store, { error: 'Failed to fetch about content', loading: false });
      }
    },
  }))
);
