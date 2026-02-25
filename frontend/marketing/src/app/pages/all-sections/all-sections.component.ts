import { Component, OnInit, AfterViewInit, OnDestroy, inject, NgZone, signal } from '@angular/core';
import { CommonModule, Location } from '@angular/common';
import { Router, NavigationEnd } from '@angular/router';
import { Title } from '@angular/platform-browser';
import { filter } from 'rxjs/operators';
import { Subscription } from 'rxjs';
import { HomeComponent } from '../home/home.component';
import { AboutComponent } from '../about/about.component';
import { GrooveJrComponent } from '../groove-jr/groove-jr.component';
import { OldSiteComponent } from '../old-site/old-site.component';
import { BlogComponent } from '../blog/blog.component';
import { FocusComponent } from '../focus/focus.component';
import { LatestPostsComponent } from '../latest-posts/latest-posts.component';
import { ScrollFadeInDirective } from '../../directives/scroll-fade-in.directive';

@Component({
  selector: 'app-all-sections',
  standalone: true,
  imports: [CommonModule, HomeComponent, AboutComponent, GrooveJrComponent, OldSiteComponent, BlogComponent, FocusComponent, LatestPostsComponent, ScrollFadeInDirective],
  templateUrl: './all-sections.component.html',
})
export class AllSectionsComponent implements OnInit, AfterViewInit, OnDestroy {
  private readonly router = inject(Router);
  private readonly location = inject(Location);
  private readonly ngZone = inject(NgZone);
  private readonly titleService = inject(Title);
  private routerSubscription?: Subscription;
  private observer?: IntersectionObserver;
  private isAutoScrolling = false;
  private lastNavTime = 0;
  
  public readonly activeSection = signal<string>('home');

  ngOnInit() {
    this.routerSubscription = this.router.events.pipe(
      filter(event => event instanceof NavigationEnd)
    ).subscribe(() => {
      this.isAutoScrolling = true;
      this.lastNavTime = Date.now();
      
      // Small delay to allow scroll restoration to complete first
      setTimeout(() => {
        this.scrollToSection();
      }, 100);

      // Allow scroll tracking after a delay
      setTimeout(() => {
        this.isAutoScrolling = false;
      }, 2500);
    });
  }

  ngAfterViewInit() {
    this.setupIntersectionObserver();
    // Small delay to ensure DOM is ready
    setTimeout(() => this.scrollToSection(), 100);
  }

  ngOnDestroy() {
    this.routerSubscription?.unsubscribe();
    this.observer?.disconnect();
  }

  private setupIntersectionObserver() {
    const options = {
      root: null,
      rootMargin: '-20% 0px -70% 0px', // Trigger when section is in the top portion of the viewport
      threshold: 0
    };

    this.observer = new IntersectionObserver((entries) => {
      if (this.isAutoScrolling || (Date.now() - this.lastNavTime < 2500)) return;

      entries.forEach(entry => {
        if (entry.isIntersecting) {
          const sectionId = entry.target.id;
          const path = sectionId === 'home' ? '' : sectionId;
          const currentPath = this.location.path().replace(/^\//, '');
          
          if (currentPath !== path) {
            this.ngZone.run(() => {
              this.activeSection.set(sectionId);
              this.location.replaceState(path);
              
              const titleMap: Record<string, string> = {
                '': 'Nate Schieber - Software Engineer',
                'focus': 'Focus - Nate Schieber',
                'latest-posts': 'Latest Posts - Nate Schieber',
                'about': 'About - Nate Schieber',
                'groovejr': 'Groove Jr. - Nate Schieber',
                'old-site-preview': 'Old Site - Nate Schieber',
                'blog': 'Blog - Nate Schieber'
              };
              const newTitle = titleMap[path] || 'Nate Schieber - Software Engineer';
              this.titleService.setTitle(newTitle);
            });
          }
        }
      });
    }, options);

    const sections = ['home', 'focus', 'latest-posts', 'about', 'groovejr', 'old-site-preview', 'blog'];
    sections.forEach(id => {
      const element = document.getElementById(id);
      if (element) {
        this.observer?.observe(element);
      }
    });
  }

  private scrollToSection() {
    const url = this.router.url;
    let sectionId = 'home';
    
    if (url.includes('focus')) sectionId = 'focus';
    else if (url.includes('latest-posts')) sectionId = 'latest-posts';
    else if (url.includes('about')) sectionId = 'about';
    else if (url.includes('groovejr')) sectionId = 'groovejr';
    else if (url.includes('old-site-preview')) sectionId = 'old-site-preview';
    else if (url.includes('blog')) sectionId = 'blog';

    this.activeSection.set(sectionId);
    this._scrollToSection(sectionId);
  }

  _scrollToSection(elementId: string) {
    const element = document.getElementById(elementId);
    if (!element)
      return;

    element.scrollIntoView({ behavior: 'smooth', block: 'start', inline: 'nearest' });
  }
}
