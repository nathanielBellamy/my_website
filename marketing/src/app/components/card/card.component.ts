import { Component, input } from '@angular/core';
import { CommonModule } from '@angular/common';
import { MarkdownComponent } from 'ngx-markdown';

@Component({
  selector: 'app-card',
  standalone: true,
  imports: [CommonModule, MarkdownComponent],
  templateUrl: './card.component.html',
})
export class CardComponent {
  title = input.required<string>();
  content = input.required<string>();
  tags = input<string[]>();
}
