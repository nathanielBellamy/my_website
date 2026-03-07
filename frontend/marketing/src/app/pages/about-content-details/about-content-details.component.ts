import { Component, OnInit, inject, signal } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ActivatedRoute, RouterLink } from '@angular/router';
import { MarkdownComponent } from 'ngx-markdown';
import { Title, Meta } from '@angular/platform-browser';
import { AboutService } from '../../services/about.service';
import { AboutContent } from '../../models/about.model';
import { ScrollFadeInDirective } from '../../directives/scroll-fade-in.directive';

@Component({
  selector: 'app-about-content-details',
  standalone: true,
  imports: [CommonModule, MarkdownComponent, RouterLink, ScrollFadeInDirective],
  templateUrl: './about-content-details.component.html',
})
export class AboutContentDetailsComponent implements OnInit {
  private readonly route = inject(ActivatedRoute);
  private readonly aboutService = inject(AboutService);
  private readonly titleService = inject(Title);
  private readonly metaService = inject(Meta);

  aboutContent = signal<AboutContent | null>(null);
  loading = signal(true);
  error = signal<string | null>(null);

  async ngOnInit() {
    const id = this.route.snapshot.paramMap.get('id');
    if (!id) {
      this.error.set('No about content ID provided');
      this.loading.set(false);
      return;
    }

    try {
      const content = await this.aboutService.getById(id);
      this.aboutContent.set(content);
      this.titleService.setTitle(`${content.title} - Nate Schieber`);
      const description = content.content.substring(0, 150).replace(/[#*`]/g, '') + '...';
      this.metaService.updateTag({ name: 'description', content: description });
    } catch (err) {
      this.error.set('Failed to load about content');
    } finally {
      this.loading.set(false);
    }
  }
}
