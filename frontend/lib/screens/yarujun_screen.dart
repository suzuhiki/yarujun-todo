import 'package:flutter/material.dart';

class YarujunScreen extends StatefulWidget {
  const YarujunScreen({Key? key}) : super(key: key);

  @override
  State<YarujunScreen> createState() => _YarujunScreenState();
}

class _YarujunScreenState extends State<YarujunScreen> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('やる順 Todo'),
        automaticallyImplyLeading: false,
      ),
      body: const Center(
        child: Text('やる順 Todo'),
      ),
    );
  }
}
