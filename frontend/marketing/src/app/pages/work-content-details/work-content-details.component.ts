import { Component, OnInit, inject, signal } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ActivatedRoute, RouterLink } from '@angular/router';
import { MarkdownComponent } from 'ngx-markdown';
import { Title, Meta } from '@angular/platform-browser';
import { WorkService } from '../../services/work.service';
import { WorkContent } from '../../models/work.model';
import { ScrollFadeInDirective } from '../../directives/scroll-fade-in.directive';
import { decodeId } from '../../utils/id-encoder';

@Component({
  selector: 'app-work-content-details',
  standalone: true,
  imports: [CommonModule, MarkdownComponent, RouterLink, ScrollFadeInDirective],
  templateUrl: './work-content-details.component.html',
})
export class WorkContentDetailsComponent implements OnInit {
  private readonly route = inject(ActivatedRoute);
  private readonly workService = inject(WorkService);
  private readonly titleService = inject(Title);
  private readonly metaService = inject(Meta);

  workContent = signal<WorkContent | null>(null);
  loading = signal(true);
  error = signal<string | null>(null);

  async ngOnInit() {
    const rawId = this.route.snapshot.paramMap.get('id');
    if (!rawId) {
      this.error.set('No work content ID provided');
      this.loading.set(false);
      return;
    }

    const id = decodeId(rawId);

    try {
      const content = await this.workService.getById(id);
      this.workContent.set(content);
      this.titleService.setTitle(`${content.title} - Nate Schieber`);
      const description = content.content.substring(0, 150).replace(/[#*`]/g, '') + '...';
      this.metaService.updateTag({ name: 'description', content: description });
    } catch (err) {
      this.error.set('Failed to load work content');
    } finally {
      this.loading.set(false);
    }
  }
}
