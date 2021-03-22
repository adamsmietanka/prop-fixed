# Introduction
AWS Lambda function written in Go calculating efficiency of fixed pitch propellers

# How does it work?
All the necessarry data is parsed from the query string of the GET method:
```
?max_speed=150&step_size=10&diameter=3.902&blades=3&cp=0.0902&prop_speed=20&power=800&angle=30&ratio=0.4&angle=30
```

The program loads meshes from the ```\data``` directory based on the number of propeller blades. The length of points array is fixed.
The first 10 points lie in the operational range of the engine rpm. 
The last five points are capped at the speed limit as the engine can't spin faster than that.
A combination of linear and binary search is used to find the triangles each point lies on. The mesh is rectangular on the XY plane, so it makes things easier. 
Barycentric coordinates are then used to calculate the desired coordinate.

## Results
The operational range of the engine can be seen as the blue dots on the chart.

<p align=middle> 
  <img src="https://github.com/adamsmietanka/prop-fixed/blob/master/docs/cp_chart.png" />
</p>

The points exceeding speed limit of the engine are displayed as red markers.

<p align=middle> 
  <img src="https://github.com/adamsmietanka/prop-fixed/blob/master/docs/eff_chart.png" />
</p>

The interpolation functions lie in a file outside ```main.go``` called ```interpolate.go```. The functions are thoroughly tested thanks to the use of TDD.

## Impact
The serverless approach replaced the original Flask backend and helped speed up the calculation process almost 40x times.
It also reduced the cost to close to nothing as there is no server nor database left to maintain.
