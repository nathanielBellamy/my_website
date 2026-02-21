import { Component } from '@angular/core';
import { ScrollFadeInDirective } from '../../directives/scroll-fade-in.directive';
import { ScrollIndicatorComponent } from '../../shared/components/scroll-indicator/scroll-indicator.component';

@Component({
  selector: 'app-focus',
  standalone: true,
  imports: [ScrollFadeInDirective, ScrollIndicatorComponent],
  templateUrl: './focus.component.html',
})
export class FocusComponent {
}