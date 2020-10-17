import { LandingPage } from '../pages/Landing';
import { LobbyPage } from '../pages/Lobby';
// import { GamePage } from '../pages/Game';
import { Pathname } from './types';

export const pages = [
  {
    route: {
      path: Pathname.LANDING,
      component: LandingPage,
      exact: true,
    },
  },
  {
    route: {
      path: Pathname.LOBBY,
      component: LobbyPage,
      exact: true,
    },
  },
  // {
  //   route: {
  //     path: Pathname.GAME,
  //     component: GamePage,
  //     exact: true,
  //   },
  // },
];
