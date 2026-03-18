import { render, screen } from '@testing-library/angular';
import { AppComponent } from './app.component';
import { Component } from '@angular/core';
import { provideRouter } from '@angular/router';

@Component({ selector: 'app-dummy', template: '<p>dummy</p>' })
class DummyComponent {}

describe('AppComponent', () => {
  it('should render the router outlet', async () => {
    await render(AppComponent, {
      providers: [
        provideRouter([{ path: '**', component: DummyComponent }]),
      ],
    });

    expect(screen.getByText('dummy')).toBeInTheDocument();
  });
});

  