import { ComponentFixture, TestBed } from '@angular/core/testing';
import { EditGrooveJrContentComponent } from './edit-groove-jr-content.component';
import { GrooveJrService } from '../../services/groove-jr.service';
import { ActivatedRoute, Router } from '@angular/router';
import { FormsModule } from '@angular/forms';
import { of } from 'rxjs';
import { GrooveJrContent } from '../../models/data-models';

describe('EditGrooveJrContentComponent', () => {
  let component: EditGrooveJrContentComponent;
  let fixture: ComponentFixture<EditGrooveJrContentComponent>;
  let mockGrooveJrService: Partial<GrooveJrService>;
  let mockActivatedRoute: Partial<ActivatedRoute>;
  let mockRouter: Partial<Router>;

  const mockGrooveJrContent: GrooveJrContent = { id: '1', title: 'Existing GrooveJr', content: 'Existing Content' };

  beforeEach(async () => {
    mockGrooveJrService = {
      getGrooveJrContentById: jasmine.createSpy('getGrooveJrContentById').and.returnValue(Promise.resolve(mockGrooveJrContent)),
      updateGrooveJrContent: jasmine.createSpy('updateGrooveJrContent').and.returnValue(Promise.resolve(mockGrooveJrContent)),
    };
    mockActivatedRoute = {
      paramMap: of(new Map([['id', '1']])),
    };
    mockRouter = {
      navigate: jasmine.createSpy('navigate'),
    };

    await TestBed.configureTestingModule({
      imports: [EditGrooveJrContentComponent, FormsModule],
      providers: [
        { provide: GrooveJrService, useValue: mockGrooveJrService },
        { provide: ActivatedRoute, useValue: mockActivatedRoute },
        { provide: Router, useValue: mockRouter },
      ],
    }).compileComponents();

    fixture = TestBed.createComponent(EditGrooveJrContentComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should fetch GrooveJr content on init and populate form', async () => {
    await fixture.whenStable();

    expect(mockGrooveJrService.getGrooveJrContentById).toHaveBeenCalledWith('1');
    expect(component.grooveJrContent()).toEqual(mockGrooveJrContent);
  });

  it('should update GrooveJr content and navigate on success', async () => {
    await fixture.whenStable();

    const updatedTitle = 'Updated Title';
    const updatedContent = 'Updated Content';
    component.grooveJrContent.set({ ...mockGrooveJrContent, title: updatedTitle, content: updatedContent });

    await component.updateContent();

    expect(mockGrooveJrService.updateGrooveJrContent).toHaveBeenCalledWith({ ...mockGrooveJrContent, title: updatedTitle, content: updatedContent });
    expect(mockRouter.navigate).toHaveBeenCalledWith(['/groovejr']);
  });

  it('should navigate back to list on goBack', () => {
    component.goBack();
    expect(mockRouter.navigate).toHaveBeenCalledWith(['/groovejr']);
  });
});
