import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Component } from '@angular/core';
import { CsvControlsComponent } from './csv-controls.component';
import { CsvService } from '../../services/csv.service';
import { of } from 'rxjs';
import { HttpResponse } from '@angular/common/http';
import { By } from '@angular/platform-browser';

@Component({
  template: `<app-csv-controls [entity]="entity"></app-csv-controls>`,
  standalone: true,
  imports: [CsvControlsComponent]
})
class TestHostComponent {
  entity = 'test-entity';
}

describe('CsvControlsComponent', () => {
  let fixture: ComponentFixture<TestHostComponent>;
  let csvServiceMock: any;

  beforeEach(async () => {
    csvServiceMock = {
      exportCsv: jest.fn(),
      importCsv: jest.fn()
    };

    await TestBed.configureTestingModule({
      imports: [TestHostComponent, CsvControlsComponent],
      providers: [
        { provide: CsvService, useValue: csvServiceMock }
      ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(TestHostComponent);
    fixture.detectChanges();
  });

  afterEach(() => {
    jest.restoreAllMocks();
  });

  it('should create', () => {
    const component = fixture.debugElement.query(By.directive(CsvControlsComponent)).componentInstance;
    expect(component).toBeTruthy();
  });

  it('should call exportCsv on export button click', () => {
    const mockBlob = new Blob([''], { type: 'text/csv' });
    const mockResponse = new HttpResponse({ body: mockBlob });
    csvServiceMock.exportCsv.mockReturnValue(of(mockResponse));

    window.URL.createObjectURL = jest.fn().mockReturnValue('blob:url');
    window.URL.revokeObjectURL = jest.fn();
    
    // Create spy for anchor click
    const clickSpy = jest.fn();
    const anchorMock = {
      click: clickSpy,
      href: '',
      download: '',
      style: {},
      setAttribute: jest.fn(), // Add setAttribute just in case
      appendChild: jest.fn()
    } as unknown as HTMLAnchorElement;

    const originalCreateElement = document.createElement.bind(document);
    jest.spyOn(document, 'createElement').mockImplementation((tagName: string, options?: any) => {
        if (tagName === 'a') {
            return anchorMock;
        }
        return originalCreateElement(tagName, options);
    });

    const component = fixture.debugElement.query(By.directive(CsvControlsComponent)).componentInstance;
    component.onExport();

    expect(csvServiceMock.exportCsv).toHaveBeenCalledWith('test-entity');
    expect(window.URL.createObjectURL).toHaveBeenCalledWith(mockBlob);
    expect(clickSpy).toHaveBeenCalled();
  });

  it('should call importCsv on file selection', () => {
    const file = new File(['test'], 'test.csv', { type: 'text/csv' });
    const event = { target: { files: [file], value: 'fakepath' } } as unknown as Event;
    
    csvServiceMock.importCsv.mockReturnValue(of({}));
    jest.spyOn(window, 'alert').mockImplementation(() => {});

    const component = fixture.debugElement.query(By.directive(CsvControlsComponent)).componentInstance;
    component.onFileSelected(event);

    expect(csvServiceMock.importCsv).toHaveBeenCalledWith('test-entity', file);
    expect(window.alert).toHaveBeenCalledWith('Import successful');
  });
});
