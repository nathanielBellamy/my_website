import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterLink } from '@angular/router';
import { ScrollFadeInDirective } from '../../directives/scroll-fade-in.directive';

@Component({
  selector: 'app-home',
  standalone: true,
  imports: [CommonModule, RouterLink, ScrollFadeInDirective],
  templateUrl: './home.component.html',
})
export class HomeComponent {
}
