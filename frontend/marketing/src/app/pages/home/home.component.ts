import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterLink } from '@angular/router';
import { ScrollFadeInDirective } from '../../directives/scroll-fade-in.directive';
import { PageComponent } from '../../components/page/page.component';

@Component({
  selector: 'app-home',
  standalone: true,
  imports: [CommonModule, RouterLink, ScrollFadeInDirective, PageComponent],
  templateUrl: './home.component.html',
})
export class HomeComponent {
}
