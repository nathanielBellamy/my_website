import { Component, inject, OnInit, signal } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { GrooveJrService } from '../../../services/groove-jr.service';
import { ActivatedRoute, Router } from '@angular/router';
import { GrooveJrContent } from '../../../models/data-models';

@Component({
  selector: 'app-edit-groove-jr-content',
  standalone: true,
  imports: [FormsModule],
  templateUrl: './edit-groove-jr-content.component.html',
  styleUrl: './edit-groove-jr-content.component.css',
})
export class EditGrooveJrContentComponent implements OnInit {
  private readonly grooveJrService = inject(GrooveJrService);
  private readonly route = inject(ActivatedRoute);
  private readonly router = inject(Router);

  grooveJrContent = signal<GrooveJrContent | undefined>(undefined);

  ngOnInit() {
    this.route.paramMap.subscribe(params => {
      const id = params.get('id');
      if (id) {
        this.grooveJrService.getGrooveJrContentById(id).then(content => {
          this.grooveJrContent.set(content);
        }).catch(error => {
          console.error('Error fetching GrooveJr content:', error);
        });
      }
    });
  }

  async updateContent() {
    if (this.grooveJrContent()) {
      try {
        await this.grooveJrService.updateGrooveJrContent(this.grooveJrContent() as GrooveJrContent);
        this.router.navigate(['/groovejr']); // Navigate back to the list after update
      } catch (error) {
        console.error('Error updating GrooveJr content:', error);
      }
    }
  }

  goBack() {
    this.router.navigate(['/groovejr']);
  }
}
