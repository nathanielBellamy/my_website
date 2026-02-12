import { render, screen } from '@testing-library/angular';
import { AppComponent } from './app.component';
import { NavbarComponent } from './components/navbar/navbar.component';
import { HeaderComponent } from './components/header/header.component';
import { RouterTestingModule } from '@angular/router/testing';
import { Component } from '@angular/core';

@Component({ template: '' })
class DummyComponent {}

describe('AppComponent', () => {

    it('should render the navbar', async () => {

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

        providers: [],

      });

      screen.getByRole('navigation');

    });

  });

  