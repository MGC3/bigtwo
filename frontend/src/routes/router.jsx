import React from 'react';
import { Switch, Route } from 'react-router-dom';
import { pages } from './routes';

export const Routes = () => {
  return (
    <Switch>
      {/* Routes */}
      {pages.map(({ route }) => {
        if ('component' in route) {
          return <Route key={route.path} {...route} />;
        }
      })}
    </Switch>
  );
};
