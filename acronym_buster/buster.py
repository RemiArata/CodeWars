def acronym_buster(message):
    known = {
        "KPI": "key performance indicators",
        'EOD': "the end of the day" ,
        'TBD': "to be decided",
        'WAH': "work at home",
        'IAM': "in a meeting",
        'OOO': "out of office",
        'NRN': "no reply necessary",
        'CTA': "call to action",
        'SWOT': "strengths, weaknesses, opportunities and threats",
    }
    sentences = message.split(".")
    for i, s in enumerate(sentences):
        msg = s.split(" ")
        for j, wrd in enumerate(msg):
            if len(wrd) >= 3 and wrd.isupper():
                if wrd in known.keys():
                    msg[j] = known[wrd]
                else:
                    return f"{wrd} is an acronym. I do not like acronyms. Please remove them from your email."
        new = " ".join(msg)
        if len(new) > 0:
            if new[0] == " ":
                new = new[1:]
            if new.isupper() or new[0].isupper():
                sentences[i] = new
            else:    
                sentences[i] = new[0].capitalize() + new[1:]
    dat_new_new = ". ".join(sentences)
    return dat_new_new if dat_new_new[-1] != " " else dat_new_new[:-1]


# print(acronym_buster("WAH"))
# print(acronym_buster("PB"))
# print(acronym_buster("HATDBEA"))
# print(acronym_buster("Going to WAH today. NRN. OOO"))
# print(acronym_buster("My SM account needs some work."))
print(acronym_buster("'WPx. rivNl XvNy. ABSdNbgNFk. SWOT. CgImUQD TBD Lr OOO CTA NRN. OOO. HQXd. WAH. Yuzx. KPI. EcaddmMGIZ.'"))

# https://www.codewars.com/kata/58397ee871df657929000209/train/python


# 'WPx. Rivnl xvny. ABSdNbgNFk. Strengths, weaknesses, opportunities and threats. CgImUQD to be decided Lr out of office call to action no reply necessary. Out of office. HQXd. Work at home. Yuzx. Key performance indicators. EcaddmMGIZ.'
# 'WPx. RivNl XvNy. ABSdNbgNFk. Strengths, weaknesses, opportunities and threats. CgImUQD to be decided Lr out of office call to action no reply necessary. Out of office. HQXd. Work at home. Yuzx. Key performance indicators. EcaddmMGIZ.'