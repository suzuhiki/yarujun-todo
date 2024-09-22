import 'package:flutter/material.dart';
import 'package:frontend/screens/home_screens.dart';
import 'package:google_fonts/google_fonts.dart';
import 'package:flutter_localizations/flutter_localizations.dart';

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
      localizationsDelegates: const [
        GlobalMaterialLocalizations.delegate,
        GlobalWidgetsLocalizations.delegate,
        GlobalCupertinoLocalizations.delegate,
      ],
      supportedLocales: const [
        Locale("ja", "JP"),
      ],
    );
  }
}
