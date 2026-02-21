import { Component, input } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterLink } from '@angular/router';
import { ScrollFadeInDirective } from '../../../directives/scroll-fade-in.directive';

@Component({
  selector: 'app-scroll-indicator',
  standalone: true,
  imports: [CommonModule, RouterLink, ScrollFadeInDirective],
  templateUrl: './scroll-indicator.component.html',
})
export class ScrollIndicatorComponent {
  target = input.required<string>();
}
