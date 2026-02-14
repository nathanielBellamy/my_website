import { Component, inject, OnInit, signal } from '@angular/core';
import { GrooveJrService } from '../../../services/groove-jr.service';
import { GrooveJrContent } from '../../../models/data-models';
import { RouterLink } from '@angular/router';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-groove-jr-content-list',
  standalone: true,
  imports: [RouterLink, CommonModule],
  templateUrl: './groove-jr-content-list.component.html',
  styleUrl: './groove-jr-content-list.component.css',
})
export class GrooveJrContentListComponent implements OnInit {
  private readonly grooveJrService = inject(GrooveJrService);
  grooveJrContent = signal<GrooveJrContent[]>([]);

  ngOnInit() {
    this.fetchGrooveJrContent();
  }

  fetchGrooveJrContent() {
    this.grooveJrService.getAllGrooveJrContent().then((content) => {
      this.grooveJrContent.set(content);
    }).catch((error) => {
      console.error('Error fetching GrooveJr content:', error);
    });
  }

  deleteContent(id: string) {
    this.grooveJrService.deleteGrooveJrContent(id).then(() => {
      this.fetchGrooveJrContent(); // Refresh the list after deletion
    }).catch((error) => {
      console.error('Error deleting GrooveJr content:', error);
    });
  }
}
