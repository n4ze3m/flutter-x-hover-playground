import 'package:flutter/material.dart';
import 'package:flutter/services.dart';

void main() {
  runApp(const MyApp());
}
const platformChannelBattery =
     MethodChannel('com.myapp/battery');

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Flutter Demo',
      theme: ThemeData(
        primarySwatch: Colors.blue,
      ),
      home: const MyHomePage(title: 'Flutter X Hover'),
    );
  }
}

class MyHomePage extends StatefulWidget {
  const MyHomePage({super.key, required this.title});

  final String title;

  @override
  State<MyHomePage> createState() => _MyHomePageState();
}

class _MyHomePageState extends State<MyHomePage> {
  // call the method channel to get the battery object
  var _battery = 'Click the button to get the battery info';
  Future<void> _getBattery() async {
    try {
      // result will be json object not int
      final  result =
          await platformChannelBattery.invokeMethod('getBattery');
      setState(() {
        _battery = '$result';
      });
    } on PlatformException catch (e) {
      setState(() {
        _battery = "Failed to get battery info: '${e.message}'.";
      });
    }
  }  

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text(widget.title),
      ),
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: <Widget>[
            Text(
              '$_battery',
              style: Theme.of(context).textTheme.headline4,
            ),
          ],
        ),
      ),
      floatingActionButton: FloatingActionButton(
        onPressed: _getBattery,
        tooltip: 'Increment',
        child: const Icon(Icons.add),
      ),
    );
  }
}
