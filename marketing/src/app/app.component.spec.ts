import { render, screen } from '@testing-library/angular';
import { AppComponent } from './app.component';
import { NavbarComponent } from './components/navbar/navbar.component';
import { HeaderComponent } from './components/header/header.component';
import { TrackerService } from './services/tracker.service';
import { RouterTestingModule } from '@angular/router/testing';
import { Component } from '@angular/core';

@Component({ template: '' })
class DummyComponent {}

describe('AppComponent', () => {

  it('should render the navbar', async () => {
    const mockTrackerService = {
      trackIp: jest.fn(),
    };
    await render(AppComponent, {
      imports: [
        NavbarComponent,
        HeaderComponent,
        RouterTestingModule.withRoutes([
          { path: '', component: DummyComponent },
          { path: 'about', component: DummyComponent },
          { path: 'groovejr', component: DummyComponent },
          { path: 'blog', component: DummyComponent },
        ]),
      ],
      providers: [
        { provide: TrackerService, useValue: mockTrackerService }
      ],
    });
    screen.getByRole('navigation');
  });

  it('should render the header', async () => {
    const mockTrackerService = {
      trackIp: jest.fn(),
    };
    await render(AppComponent, {
      imports: [
        NavbarComponent,
        HeaderComponent,
        RouterTestingModule.withRoutes([
          { path: '', component: DummyComponent },
          { path: 'about', component: DummyComponent },
          { path: 'groovejr', component: DummyComponent },
          { path: 'blog', component: DummyComponent },
        ]),
      ],
      providers: [
        { provide: TrackerService, useValue: mockTrackerService }
      ],
    });
    screen.getByRole('banner');
  });

  it('should call trackIp on init', async () => {
    const mockTrackerService = {
      trackIp: jest.fn(),
    };
    await render(AppComponent, {
      imports: [
        NavbarComponent,
        HeaderComponent,
        RouterTestingModule.withRoutes([
          { path: '', component: DummyComponent },
          { path: 'about', component: DummyComponent },
          { path: 'groovejr', component: DummyComponent },
          { path: 'blog', component: DummyComponent },
        ]),
      ],
      providers: [
        { provide: TrackerService, useValue: mockTrackerService }
      ],
    });
    expect(mockTrackerService.trackIp).toHaveBeenCalled();
  });
});