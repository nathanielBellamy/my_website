import { Component, OnInit, inject, signal } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ActivatedRoute, RouterLink } from '@angular/router';
import { MarkdownComponent } from 'ngx-markdown';
import { Title, Meta } from '@angular/platform-browser';
import { HomeService } from '../../services/home.service';
import { HomeContent } from '../../models/home.model';
import { ScrollFadeInDirective } from '../../directives/scroll-fade-in.directive';
import { decodeId } from '../../utils/id-encoder';

@Component({
  selector: 'app-home-content-details',
  standalone: true,
  imports: [CommonModule, MarkdownComponent, RouterLink, ScrollFadeInDirective],
  templateUrl: './home-content-details.component.html',
})
export class HomeContentDetailsComponent implements OnInit {
  private readonly route = inject(ActivatedRoute);
  private readonly homeService = inject(HomeService);
  private readonly titleService = inject(Title);
  private readonly metaService = inject(Meta);

  homeContent = signal<HomeContent | null>(null);
  loading = signal(true);
  error = signal<string | null>(null);

  async ngOnInit() {
    const rawId = this.route.snapshot.paramMap.get('id');
    if (!rawId) {
      this.error.set('No home content ID provided');
      this.loading.set(false);
      return;
    }

    const id = decodeId(rawId);

    try {
      const content = await this.homeService.getById(id);
      this.homeContent.set(content);
      this.titleService.setTitle(`${content.title} - Nate Schieber`);
      const description = content.content.substring(0, 150).replace(/[#*`]/g, '') + '...';
      this.metaService.updateTag({ name: 'description', content: description });
    } catch (err) {
      this.error.set('Failed to load home content');
    } finally {
      this.loading.set(false);
    }
  }
}
