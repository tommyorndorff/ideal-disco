import http from 'k6/http';
import { check, sleep } from 'k6';

export const options = {
  stages: [
    { duration: '30s', target: 100 },
    // { duration: '2m', target: 35 },
    { duration: '10m', target: 100 },
    // { duration: '1m30s', target: 10 },
    { duration: '20s', target: 5 },
  ],
};

export default function() {
    let response = http.get("http://tornd-appli-189ys0tuft122-1472934720.us-east-1.elb.amazonaws.com:8080/pi/10000");
};
