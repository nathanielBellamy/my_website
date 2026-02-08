import { render, screen } from '@testing-library/angular';
import { AppComponent } from './app.component';
import { appConfig } from './app.config';

describe('AppComponent', () => {
  it('should create the app', async () => {
    await render(AppComponent, {
      imports: [],
      providers: appConfig.providers,
    });
    expect(screen.getByText('admin')).toBeInTheDocument();
  });
});
