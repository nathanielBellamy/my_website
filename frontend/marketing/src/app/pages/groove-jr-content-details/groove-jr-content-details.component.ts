import { Component, OnInit, inject, signal } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ActivatedRoute, RouterLink } from '@angular/router';
import { MarkdownComponent } from 'ngx-markdown';
import { Title, Meta } from '@angular/platform-browser';
import { GrooveJrService } from '../../services/groove-jr.service';
import { GrooveJrContent } from '../../models/groove-jr.model';
import { ScrollFadeInDirective } from '../../directives/scroll-fade-in.directive';

@Component({
  selector: 'app-groove-jr-content-details',
  standalone: true,
  imports: [CommonModule, MarkdownComponent, RouterLink, ScrollFadeInDirective],
  templateUrl: './groove-jr-content-details.component.html',
})
export class GrooveJrContentDetailsComponent implements OnInit {
  private readonly route = inject(ActivatedRoute);
  private readonly grooveJrService = inject(GrooveJrService);
  private readonly titleService = inject(Title);
  private readonly metaService = inject(Meta);

  grooveJrContent = signal<GrooveJrContent | null>(null);
  loading = signal(true);
  error = signal<string | null>(null);

  async ngOnInit() {
    const id = this.route.snapshot.paramMap.get('id');
    if (!id) {
      this.error.set('No GrooveJr content ID provided');
      this.loading.set(false);
      return;
    }

    try {
      const content = await this.grooveJrService.getById(id);
      this.grooveJrContent.set(content);
      this.titleService.setTitle(`${content.title} - Nate Schieber`);
      const description = content.content.substring(0, 150).replace(/[#*`]/g, '') + '...';
      this.metaService.updateTag({ name: 'description', content: description });
    } catch (err) {
      this.error.set('Failed to load GrooveJr content');
    } finally {
      this.loading.set(false);
    }
  }
}
