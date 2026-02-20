import { Component } from '@angular/core';
import { ScrollFadeInDirective } from '../../directives/scroll-fade-in.directive';

@Component({
  selector: 'app-focus',
  standalone: true,
  imports: [ScrollFadeInDirective],
  templateUrl: './focus.component.html',
})
export class FocusComponent {
}