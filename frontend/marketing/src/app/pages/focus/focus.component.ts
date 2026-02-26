import { Component } from '@angular/core';
import { ScrollFadeInDirective } from '../../directives/scroll-fade-in.directive';
import { PageComponent } from '../../components/page/page.component';

@Component({
  selector: 'app-focus',
  standalone: true,
  imports: [ScrollFadeInDirective, PageComponent],
  templateUrl: './focus.component.html',
})
export class FocusComponent {
  readonly technologies = [
    'NixOS', 'Vim', 'LLMs', 'TypeScript', 'Java', 'Scala', 'Go', 'C++', 'Rust', 'Ruby',
    'PostgreSQL', 'SQLite', 'Qt', 'Angular', 'Vue', 'Spring',
    'Rails', 'Akka', 'CAF', ''
  ];
}