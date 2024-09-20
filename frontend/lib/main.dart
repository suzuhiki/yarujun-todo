import 'package:flutter/material.dart';
import 'package:frontend/screens/home_screens.dart';
import 'package:frontend/screens/login_screen.dart';
import 'package:frontend/screens/tasks_screen.dart';
import 'package:google_fonts/google_fonts.dart';

void main() {
  runApp(const MainApp());
}

class MainApp extends StatelessWidget {
  const MainApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      theme: ThemeData(
        primarySwatch: Colors.blue,
        textTheme: GoogleFonts.notoSansTextTheme(Theme.of(context).textTheme),
        fontFamily: GoogleFonts.notoSans().fontFamily,
        appBarTheme: const AppBarTheme(
          backgroundColor: Colors.white30,
        ),
      ),
      home: Container(
        child: SafeArea(child: HomeScreens()),
        color: Colors.white,
      ),
    );
  }
}
