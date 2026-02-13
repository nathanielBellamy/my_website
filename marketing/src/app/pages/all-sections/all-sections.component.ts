import { Component, OnInit, AfterViewInit, inject } from '@angular/core';
import { CommonModule } from '@angular/common';
import { Router, NavigationEnd } from '@angular/router';
import { filter } from 'rxjs/operators';
import { HomeComponent } from '../home/home.component';
import { AboutComponent } from '../about/about.component';
import { GrooveJrComponent } from '../groove-jr/groove-jr.component';
import { BlogComponent } from '../blog/blog.component';

@Component({
  selector: 'app-all-sections',
  standalone: true,
  imports: [CommonModule, HomeComponent, AboutComponent, GrooveJrComponent, BlogComponent],
  templateUrl: './all-sections.component.html',
})
export class AllSectionsComponent implements OnInit, AfterViewInit {
  private readonly router = inject(Router);

  ngOnInit() {
    this.router.events.pipe(
      filter(event => event instanceof NavigationEnd)
    ).subscribe(() => {
      this.scrollToSection();
    });
  }

  ngAfterViewInit() {
    // Small delay to ensure DOM is ready
    setTimeout(() => this.scrollToSection(), 100);
  }

  private scrollToSection() {
    const url = this.router.url;
    let sectionId = 'home';
    
    // Simple matching logic
    if (url.includes('about')) sectionId = 'about';
    else if (url.includes('groovejr')) sectionId = 'groovejr';
    else if (url.includes('blog')) sectionId = 'blog';
    
    const element = document.getElementById(sectionId);
    if (element) {
      element.scrollIntoView({ behavior: 'smooth', block: 'start' });
    } else {
        // fallback to top if element not found (e.g. home)
        window.scrollTo({ top: 0, behavior: 'smooth' });
    }
  }
}
