import { ComponentFixture, TestBed } from '@angular/core/testing';
import { CreateGrooveJrContentComponent } from './create-groove-jr-content.component';
import { GrooveJrService } from '../../services/groove-jr.service';
import { Router } from '@angular/router';
import { FormsModule } from '@angular/forms';
import { GrooveJrContent } from '../../models/data-models';

describe('CreateGrooveJrContentComponent', () => {
  let component: CreateGrooveJrContentComponent;
  let fixture: ComponentFixture<CreateGrooveJrContentComponent>;
  let mockGrooveJrService: Partial<GrooveJrService>;
  let mockRouter: Partial<Router>;

  beforeEach(async () => {
    mockGrooveJrService = {
      createGrooveJrContent: jasmine.createSpy('createGrooveJrContent').and.returnValue(Promise.resolve({ id: '1', title: 'New', content: 'Test' })),
    };
    mockRouter = {
      navigate: jasmine.createSpy('navigate'),
    };

    await TestBed.configureTestingModule({
      imports: [CreateGrooveJrContentComponent, FormsModule],
      providers: [
        { provide: GrooveJrService, useValue: mockGrooveJrService },
        { provide: Router, useValue: mockRouter },
      ],
    }).compileComponents();

    fixture = TestBed.createComponent(CreateGrooveJrContentComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should create GrooveJr content and navigate on success', async () => {
    const newContent: GrooveJrContent = { id: '', title: 'Test Title', content: 'Test Content' };
    component.grooveJrContent = { ...newContent };

    await component.createContent();

    expect(mockGrooveJrService.createGrooveJrContent).toHaveBeenCalledWith(newContent);
    expect(mockRouter.navigate).toHaveBeenCalledWith(['/groovejr']);
  });

  it('should navigate back to list on goBack', () => {
    component.goBack();
    expect(mockRouter.navigate).toHaveBeenCalledWith(['/groovejr']);
  });
});
