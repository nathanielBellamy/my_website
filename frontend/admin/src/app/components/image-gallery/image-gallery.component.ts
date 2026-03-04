import { Component, inject, OnInit, signal } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { ImageService } from '../../services/image.service';
import { Image } from '../../models/data-models';

@Component({
  selector: 'app-image-gallery',
  standalone: true,
  imports: [FormsModule],
  templateUrl: './image-gallery.component.html',
  styleUrl: './image-gallery.component.css',
})
export class ImageGalleryComponent implements OnInit {
  private readonly imageService = inject(ImageService);

  images = signal<Image[]>([]);
  uploading = signal<boolean>(false);
  selectedFile: File | null = null;
  altText = '';

  ngOnInit() {
    this.fetchImages();
  }

  async fetchImages() {
    try {
      const imgs = await this.imageService.listImages();
      this.images.set(imgs);
    } catch (error) {
      console.error('Error fetching images:', error);
    }
  }

  onFileSelected(event: Event) {
    const input = event.target as HTMLInputElement;
    if (input.files && input.files.length > 0) {
      this.selectedFile = input.files[0];
    }
  }

  async uploadImage() {
    if (!this.selectedFile) return;

    this.uploading.set(true);
    try {
      await this.imageService.uploadImage(this.selectedFile, this.altText);
      this.selectedFile = null;
      this.altText = '';
      // Reset file input if needed (manual reset via data binding or viewchild would be cleaner but this is simple)
      await this.fetchImages();
    } catch (error) {
      console.error('Error uploading image:', error);
    } finally {
      this.uploading.set(false);
    }
  }

  async deleteImage(id: string) {
    if (!confirm('Are you sure you want to delete this image?')) return;

    try {
      await this.imageService.deleteImage(id);
      await this.fetchImages();
    } catch (error) {
      console.error('Error deleting image:', error);
    }
  }

  copyMarkdown(image: Image) {
    const markdown = `![${image.altText || image.originalName}](/api/images/${image.filename})`;
    navigator.clipboard.writeText(markdown).then(() => {
      alert('Markdown copied to clipboard!');
    }).catch(err => {
      console.error('Could not copy text: ', err);
    });
  }
}
