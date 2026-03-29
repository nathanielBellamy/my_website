import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ScrollFadeInDirective } from '../../directives/scroll-fade-in.directive';

@Component({
  selector: 'app-privacy-policy',
  standalone: true,
  imports: [CommonModule, ScrollFadeInDirective],
  templateUrl: './privacy-policy.component.html',
})
export class PrivacyPolicyComponent {}
