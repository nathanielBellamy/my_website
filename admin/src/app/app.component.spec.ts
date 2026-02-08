import { render, screen } from '@testing-library/angular';
import { AppComponent } from './app.component';
import { appConfig } from './app.config';
import { Router } from '@angular/router';
import userEvent from '@testing-library/user-event';
import { LocationStrategy } from '@angular/common';

describe('AppComponent', () => {
  let routerSpy: any;
  let locationStrategySpy: any;

  beforeEach(async () => {
    // Create a spy object for the Router
    routerSpy = {
      navigate: jest.fn().mockResolvedValue(true),
      navigateByUrl: jest.fn().mockResolvedValue(true),
      url: '/',
      events: {
        subscribe: jest.fn(),
      },
      routerState: {
        root: {
          snapshot: {
            url: [],
            params: {},
            queryParams: {},
            fragment: null,
            data: {},
            children: [],
            firstChild: null,
          },
        },
      },
      createUrlTree: jest.fn(),
      serializeUrl: jest.fn(),
    };

    // Create a spy object for LocationStrategy
    locationStrategySpy = {
      path: jest.fn(() => ''),
      prepareExternalUrl: jest.fn((internal: string) => internal),
      pushState: jest.fn(),
      replaceState: jest.fn(),
      onPopState: jest.fn(),
      getBaseHref: jest.fn(() => ''),
    };


    jest.spyOn(routerSpy, 'navigate');
    await render(AppComponent, {
      providers: [
        { provide: Router, useValue: routerSpy },
        { provide: LocationStrategy, useValue: locationStrategySpy },
      ],
    });
  });

  it('should create the app', () => {
    expect(screen.getByText('Admin Panel')).toBeInTheDocument();
  });

  it('should display navigation links', () => {
    const homeLink = screen.getByText('Home');
    expect(homeLink).toBeInTheDocument();
    expect(homeLink).toHaveAttribute('routerLink', '/home');

    const grooveJrLink = screen.getByText('GrooveJr');
    expect(grooveJrLink).toBeInTheDocument();
    expect(grooveJrLink).toHaveAttribute('routerLink', '/groovejr');

    const aboutLink = screen.getByText('About');
    expect(aboutLink).toBeInTheDocument();
    expect(aboutLink).toHaveAttribute('routerLink', '/about');

    const blogLink = screen.getByText('Blog');
    expect(blogLink).toBeInTheDocument();
    expect(blogLink).toHaveAttribute('routerLink', '/blog');
  });

  // TODO: fix test router implementation, it's always a pain

  // it('should navigate to home when Home link is clicked', async () => {
  //   const homeLink = screen.getByText('Home');
  //   await userEvent.click(homeLink);
  //   expect(routerSpy.navigate).toHaveBeenCalledWith(['/home']);
  // });

  // it('should navigate to groovejr when GrooveJr link is clicked', async () => {
  //   const grooveJrLink = screen.getByText('GrooveJr');
  //   await userEvent.click(grooveJrLink);
  //   expect(routerSpy.navigate).toHaveBeenCalledWith(['/groovejr']);
  // });

  // it('should navigate to about when About link is clicked', async () => {
  //   const aboutLink = screen.getByText('About');
  //   await userEvent.click(aboutLink);
  //   expect(routerSpy.navigate).toHaveBeenCalledWith(['/about']);
  // });

  // it('should navigate to blog when Blog link is clicked', async () => {
  //   const blogLink = screen.getByText('Blog');
  //   await userEvent.click(blogLink);
  //   expect(routerSpy.navigate).toHaveBeenCalledWith(['/blog']);
  // });
});

