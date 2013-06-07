#! /bin/bash

# python https://developers.google.com/edu/python/
lists=("http://www.youtube.com/watch?v=tKTZoB2Vjuk"
       "http://www.youtube.com/watch?v=EPYupizJYQI"
       "http://www.youtube.com/watch?v=haycL41dAhg"
       "http://www.youtube.com/watch?v=kWyoYtvJpe4"
       "http://www.youtube.com/watch?v=uKZ8GBKmeDM"
       "http://www.youtube.com/watch?v=Nn2KQmVF5Og"
       "http://www.youtube.com/watch?v=IcteAbMC1Ok");

for l in "${lists[@]}"
do
    echo "downloading $l"
    youtube-dl "$l" -f 22/35/34
done
