# Create Markdown Table of content
Create Markdown table of content from directory structure.

## What does it do?

It creates tree like strucure for markdown based on contenet of directory

Features implemented:
1.  Choose any operator for points
    ex: -, *, etc
    default is `-`
2.  Choose indentation spacing
    default is `4`
3.  Ignore any directories or files
    Space separated when multiple values are present.
4.  Use the same for remote repositories
    Not yet implemented.

## Setup

1.  Clone this repository
    >   git clone https://github.com/ashwahegde/markdown-tree-generator.git
2.  Install Go
    Refer [docs](https://golang.org/doc/install) for help.

## How to Use It?

1.  With Default options
    ```bash
    go run go/parser.go ~/myrepos/coursera_certficates/
    ```
2.  Specify options
    ```bash
    go run go/parser.go ~/myrepos/coursera_certficates/ "*" 4 combined
    ```
    in this case:
    -   It traverses through above directory (~/myrepos/coursera_certficates/).
    -   Uses `*` for markdown points.
    -   Adds 4 spaces for indentation.
    -   Contents of `combined` directory is ignored

3.  Use it for parsing GitHub repositories
    ```bash
    go run go/parser.go https://github.com/ashwahegde/my_coursera_certficates "-" 4 combined
    ```
    in this case:
    -   It traverses through above directory (~/myrepos/coursera_certficates/).
    -   Uses `-` for markdown points.
    -   Adds 4 spaces for indentation.
    -   Contents of `combined` directory is ignored.

## Sample Output

Output looks like this:
```md
-   [business](/business)
    -   [Budgeting for Business.pdf](</business/Budgeting for Business.pdf>)
-   [data science](</data science>)
    -   [Introduction to Data Science in Python.pdf](</data science/Introduction to Data Science in Python.pdf>)
    -   [advanced ibm](</data science/advanced ibm>)
        -   [Advanced Machine Learning and Signal Processing.pdf](</data science/advanced ibm/Advanced Machine Learning and Signal Processing.pdf>)
        -   [Applied AI with DeepLearning.pdf](</data science/advanced ibm/Applied AI with DeepLearning.pdf>)
        -   [Fundamentals of Scalable Data Science.pdf](</data science/advanced ibm/Fundamentals of Scalable Data Science.pdf>)
-   [deeplearning](/deeplearning)
    -   [Machine Learning.pdf](</deeplearning/Machine Learning.pdf>)
    -   [deeplearning_ai](/deeplearning/deeplearning_ai)
        -   [Convolutional Neural Networks.pdf](</deeplearning/deeplearning_ai/Convolutional Neural Networks.pdf>)
        -   [Deep Learning Specialization.pdf](</deeplearning/deeplearning_ai/Deep Learning Specialization.pdf>)
        -   [Improving Deep Neural Networks.pdf](</deeplearning/deeplearning_ai/Improving Deep Neural Networks.pdf>)
        -   [Neural Networks and Deep Learning.pdf](</deeplearning/deeplearning_ai/Neural Networks and Deep Learning.pdf>)
        -   [Sequence Models.pdf](</deeplearning/deeplearning_ai/Sequence Models.pdf>)
        -   [Structuring Machine Learning Projects.pdf](</deeplearning/deeplearning_ai/Structuring Machine Learning Projects.pdf>)
-   [devops](/devops)
    -   [Udemy Docker.pdf](</devops/Udemy Docker.pdf>)
    -   [Udemy Git Complete.pdf](</devops/Udemy Git Complete.pdf>)
-   [iot](/iot)
    -   [An Introduction to Programming the Internet of Things (IOT).pdf](</iot/An Introduction to Programming the Internet of Things (IOT).pdf>)
    -   [Interfacing with the Arduino.pdf](</iot/Interfacing with the Arduino.pdf>)
    -   [Interfacing with the Raspberry Pi.pdf](</iot/Interfacing with the Raspberry Pi.pdf>)
    -   [Introduction to the Internet of Things and Embedded Systems.pdf](</iot/Introduction to the Internet of Things and Embedded Systems.pdf>)
    -   [Programming for the Internet of Things Project.pdf](</iot/Programming for the Internet of Things Project.pdf>)
    -   [The Arduino Platform and C Programming.pdf](</iot/The Arduino Platform and C Programming.pdf>)
    -   [The Raspberry Pi Platform and Python Programming for the Raspberry Pi.pdf](</iot/The Raspberry Pi Platform and Python Programming for the Raspberry Pi.pdf>)
-   [linux](/linux)
    -   [Linux Tools for Developers.pdf](</linux/Linux Tools for Developers.pdf>)
    -   [Linux for Developers.pdf](</linux/Linux for Developers.pdf>)
    -   [Open Source Software Development Methods.pdf](</linux/Open Source Software Development Methods.pdf>)
    -   [Open Source Software Development, Linux and Git.pdf](</linux/Open Source Software Development, Linux and Git.pdf>)
    -   [Using Git for Distributed Development.pdf](</linux/Using Git for Distributed Development.pdf>)
-   [netw](/netw)
    -   [Data Communications and Network Services.pdf](</netw/Data Communications and Network Services.pdf>)
    -   [Network Protocols and Architecture.pdf](</netw/Network Protocols and Architecture.pdf>)
-   [others](/others)
    -   [Cisco Certified Interviewer Award.pdf](</others/Cisco Certified Interviewer Award.pdf>)
-   [programming](/programming)
    -   [Data Structures.pdf](</programming/Data Structures.pdf>)
    -   [dbms](/programming/dbms)
        -   [Databases and SQL for Data Science.pdf](</programming/dbms/Databases and SQL for Data Science.pdf>)
    -   [go](/programming/go)
        -   [Go- The Complete Developer's Guide (Golang).pdf](</programming/go/Go- The Complete Developer's Guide (Golang).pdf>)
    -   [julia](/programming/julia)
        -   [Julia Scientific Programming.pdf](</programming/julia/Julia Scientific Programming.pdf>)
    -   [python](/programming/python)
        -   [Programming for Everybody.pdf](</programming/python/Programming for Everybody.pdf>)
        -   [Python Data Structures.pdf](</programming/python/Python Data Structures.pdf>)
        -   [Using Databases with Python.pdf](</programming/python/Using Databases with Python.pdf>)
        -   [Using Python to Access Web Data.pdf](</programming/python/Using Python to Access Web Data.pdf>)
-   [robotics](/robotics)
    -   [Aerial Robotics.pdf](</robotics/Aerial Robotics.pdf>)
-   [security](/security)
    -   [Cisco Security Ninja Green Belt.pdf](</security/Cisco Security Ninja Green Belt.pdf>)
    -   [Cyber Threats and Attack Vectors.pdf](</security/Cyber Threats and Attack Vectors.pdf>)
-   [tensorflow_specialization](/tensorflow_specialization)
    -   [Convolutional Neural Networks in TensorFlow.pdf](</tensorflow_specialization/Convolutional Neural Networks in TensorFlow.pdf>)
    -   [Introduction to Tensorflow.pdf](</tensorflow_specialization/Introduction to Tensorflow.pdf>)
-   [web](/web)
    -   [Introduction to HTML5.pdf](</web/Introduction to HTML5.pdf>)
```

Visualization looks like [this](https://github.com/ashwahegde/my_coursera_certficates#categories)