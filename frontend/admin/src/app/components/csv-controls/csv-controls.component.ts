import { Component, input, inject } from '@angular/core';
import { CsvService } from '../../services/csv.service';

@Component({
  selector: 'app-csv-controls',
  standalone: true,
  imports: [],
  templateUrl: './csv-controls.component.html'
})
export class CsvControlsComponent {
  entity = input.required<string>();
  private readonly csvService = inject(CsvService);

  onExport() {
    this.csvService.exportCsv(this.entity()).subscribe({
      next: (response) => {
        const blob = response.body;
        if (blob) {
          const url = window.URL.createObjectURL(blob);
          const a = document.createElement('a');
          a.href = url;
          a.download = `${this.entity()}.csv`;
          a.click();
          window.URL.revokeObjectURL(url);
        }
      },
      error: (err) => console.error('Export failed', err)
    });
  }

  onFileSelected(event: Event) {
    const input = event.target as HTMLInputElement;
    if (input.files?.length) {
      const file = input.files[0];
      this.csvService.importCsv(this.entity(), file).subscribe({
        next: () => {
          alert('Import successful');
          input.value = ''; // Reset input
        },
        error: (err) => {
          console.error('Import failed', err);
          alert('Import failed');
          input.value = ''; // Reset input
        }
      });
    }
  }
}
