import { Component } from '@angular/core';
import { ScrollFadeInDirective } from '../../directives/scroll-fade-in.directive';
import { ScrollIndicatorComponent } from '../../components/scroll-indicator/scroll-indicator.component';
import { PageComponent } from '../../components/page/page.component';

@Component({
  selector: 'app-focus',
  standalone: true,
  imports: [ScrollFadeInDirective, ScrollIndicatorComponent, PageComponent],
  templateUrl: './focus.component.html',
})
export class FocusComponent {
  readonly technologies = [
    'LLMs', 'C++', 'TypeScript', 'Java', 'Scala', 'Go', 'Rust', 'Ruby',
    'PostgreSQL', 'SQLite', 'NixOS', 'Qt', 'Angular', 'Vue', 'Spring',
    'Rails', 'Akka', 'CAF'
  ];
}