import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { PageComponent } from '../../components/page/page.component';
import { ScrollFadeInDirective } from '../../directives/scroll-fade-in.directive';
import { ScrollIndicatorComponent } from '../../components/scroll-indicator/scroll-indicator.component';

@Component({
  selector: 'app-old-site',
  standalone: true,
  imports: [CommonModule, PageComponent, ScrollFadeInDirective, ScrollIndicatorComponent],
  templateUrl: './old-site.component.html'
})
export class OldSiteComponent {}
