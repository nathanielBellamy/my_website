import { render, screen } from '@testing-library/angular';
import { HomeComponent } from './home.component';
import { CardComponent } from '../../components/card/card.component';

describe('HomeComponent', () => {
  it('should render multiple cards', async () => {
    await render(HomeComponent, {
      imports: [CardComponent],
    });
    expect(screen.getAllByRole('article').length).toBe(6); // Assuming 6 cards are rendered
  });
});