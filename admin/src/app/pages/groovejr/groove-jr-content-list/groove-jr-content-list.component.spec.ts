import { ComponentFixture, TestBed } from '@angular/core/testing';
import { GrooveJrContentListComponent } from './groove-jr-content-list.component';
import { GrooveJrService } from '../../services/groove-jr.service';
import { of } from 'rxjs';
import { GrooveJrContent } from '../../models/data-models';

describe('GrooveJrContentListComponent', () => {
  let component: GrooveJrContentListComponent;
  let fixture: ComponentFixture<GrooveJrContentListComponent>;
  let mockGrooveJrService: Partial<GrooveJrService>;

  const mockGrooveJrContent: GrooveJrContent[] = [
    { id: '1', title: 'GrooveJr 1', content: 'Content 1' },
    { id: '2', title: 'GrooveJr 2', content: 'Content 2' },
  ];

  beforeEach(async () => {
    mockGrooveJrService = {
      getAllGrooveJrContent: jasmine.createSpy('getAllGrooveJrContent').and.returnValue(Promise.resolve(mockGrooveJrContent)),
      deleteGrooveJrContent: jasmine.createSpy('deleteGrooveJrContent').and.returnValue(Promise.resolve()),
    };

    await TestBed.configureTestingModule({
      imports: [GrooveJrContentListComponent],
      providers: [{ provide: GrooveJrService, useValue: mockGrooveJrService }]
    })
    .compileComponents();

    fixture = TestBed.createComponent(GrooveJrContentListComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should fetch GrooveJr content on ngOnInit', async () => {
    await fixture.whenStable();

    expect(mockGrooveJrService.getAllGrooveJrContent).toHaveBeenCalled();
    expect(component.grooveJrContent()).toEqual(mockGrooveJrContent);
  });

  it('should delete GrooveJr content and refresh the list', async () => {
    await fixture.whenStable();
    expect(component.grooveJrContent()).toEqual(mockGrooveJrContent);

    component.deleteContent('1');

    await fixture.whenStable();

    expect(mockGrooveJrService.deleteGrooveJrContent).toHaveBeenCalledWith('1');
    expect(mockGrooveJrService.getAllGrooveJrContent).toHaveBeenCalledTimes(2);
  });
});
