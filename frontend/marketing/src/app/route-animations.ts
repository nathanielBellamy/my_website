import {
  trigger,
  transition,
  style,
  query,
  group,
  animate,
} from '@angular/animations';

export const fakeScrollAnimation = trigger('routeAnimations', [
  transition(':increment', [
    style({ position: 'relative' }),
    query(':enter, :leave', [
      style({
        position: 'absolute',
        top: 0,
        left: 0,
        width: '100%',
        height: '100%',
        opacity: 1
      })
    ], { optional: true }),
    query(':enter', [
      style({ transform: 'translateY(100vh)', opacity: 0 })
    ], { optional: true }),
    group([
      query(':leave', [
        animate('600ms cubic-bezier(0.25, 1, 0.5, 1)', style({ transform: 'translateY(-100vh)', opacity: 0 }))
      ], { optional: true }),
      query(':enter', [
        animate('600ms cubic-bezier(0.25, 1, 0.5, 1)', style({ transform: 'translateY(0)', opacity: 1 }))
      ], { optional: true })
    ])
  ]),
  transition(':decrement', [
    style({ position: 'relative' }),
    query(':enter, :leave', [
      style({
        position: 'absolute',
        top: 0,
        left: 0,
        width: '100%',
        height: '100%',
        opacity: 1
      })
    ], { optional: true }),
    query(':enter', [
      style({ transform: 'translateY(-100vh)', opacity: 0 })
    ], { optional: true }),
    group([
      query(':leave', [
        animate('600ms cubic-bezier(0.25, 1, 0.5, 1)', style({ transform: 'translateY(100vh)', opacity: 0 }))
      ], { optional: true }),
      query(':enter', [
        animate('600ms cubic-bezier(0.25, 1, 0.5, 1)', style({ transform: 'translateY(0)', opacity: 1 }))
      ], { optional: true })
    ])
  ])
]);
