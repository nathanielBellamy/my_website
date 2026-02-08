import { ComponentFixture, TestBed } from '@angular/core/testing';
import { AboutContentListComponent } from './about-content-list.component';
import { AboutService } from '../../services/about.service';
import { of } from 'rxjs';
import { AboutContent } from '../../models/data-models';

describe('AboutContentListComponent', () => {
  let component: AboutContentListComponent;
  let fixture: ComponentFixture<AboutContentListComponent>;
  let mockAboutService: Partial<AboutService>;

  const mockAboutContent: AboutContent[] = [
    { id: '1', title: 'About 1', content: 'Content 1' },
    { id: '2', title: 'About 2', content: 'Content 2' },
  ];

  beforeEach(async () => {
    mockAboutService = {
      getAllAboutContent: jasmine.createSpy('getAllAboutContent').and.returnValue(Promise.resolve(mockAboutContent)),
      deleteAboutContent: jasmine.createSpy('deleteAboutContent').and.returnValue(Promise.resolve()),
    };

    await TestBed.configureTestingModule({
      imports: [AboutContentListComponent],
      providers: [{ provide: AboutService, useValue: mockAboutService }]
    })
    .compileComponents();

    fixture = TestBed.createComponent(AboutContentListComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should fetch about content on ngOnInit', async () => {
    await fixture.whenStable();

    expect(mockAboutService.getAllAboutContent).toHaveBeenCalled();
    expect(component.aboutContent()).toEqual(mockAboutContent);
  });

  it('should delete about content and refresh the list', async () => {
    await fixture.whenStable();
    expect(component.aboutContent()).toEqual(mockAboutContent);

    component.deleteContent('1');

    await fixture.whenStable();

    expect(mockAboutService.deleteAboutContent).toHaveBeenCalledWith('1');
    expect(mockAboutService.getAllAboutContent).toHaveBeenCalledTimes(2);
  });
});
