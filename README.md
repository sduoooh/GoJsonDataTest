# GoJsonDataTest
A test source code with limitations.

## Context
When the data structure is stable and there are few nestings, only the data values change, which is an expected better scenario for the performance of the json decoder, so I chose to test in such a scenario.
By adding the requirement for a data conversion of the data passed upstream, I examined the time-consuming difference between direct regular parsing of the data string and first JSON decoding and then secondary assignment under the current requirements I encountered.


## Result
After many tests, when the number of tests is around 100 to 1,000, the results of the two are close, and the results fluctuate greatly; when the number becomes larger, to 10,000 or more, the regular expression parsing time and the json decoder parsing time stabilize at 2 Than about 3.
